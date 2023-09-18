package services

import (
	"context"
	"errors"
	"fmt"
	"start/models"
	"start/utils"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
Identifier can be both email or username
*/
func GetUser(ident models.UserIdentifier) (*models.User, error) {

	if user, err := getUserRedis(ident); err == nil {
		return user, nil
	} else {
		log.Err(err)
		return getUserPg(ident)
	}

}

/*
@Returns

	true if user was created on postgres and false if it was not
*/
func CreateUser(user models.User) (bool, error) {

	id, err := createUserPg(user)
	if err != nil {
		return false, err
	}

	user.Id = *id

	err = createUserRedis(user)
	if err != nil {
		return true, err
	}

	return true, nil
}

func DeleteUser(ident models.UserIdentifier) error {

	err := deleteUserPg(ident)
	if err != nil {
		return err
	}
	err = deleteUserRedis(ident)

	return err
}

func getUserRedis(ident models.UserIdentifier) (*models.User, error) {

	ctx := context.Background()

	cli, err := utils.RedisCli()
	if err != nil {
		return nil, fmt.Errorf("error creating a redis client %w", err)
	}

	defer func() {
		if err := cli.Close(); err != nil {
			log.Err(err)
		}
	}()

	key, err := scanUserRedis(ident)
	if err != nil {
		return nil, err
	}

	res, err := cli.LRange(ctx, *key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting values from redis %w", err)
	}

	id, err := strconv.Atoi(res[0])
	if err != nil {
		return nil, fmt.Errorf("error converting first value to int %w", err)
	}

	user := models.User{
		Id:       id,
		Username: res[1],
		PessName: res[2],
		UserMail: res[3],
		Password: res[4],
	}

	return &user, nil
}

func scanUserRedis(ident models.UserIdentifier) (*string, error) {

	ctx := context.Background()

	cli, err := utils.RedisCli()
	if err != nil {
		return nil, fmt.Errorf("error scanning redis user %w", err)
	}

	defer func() {
		if err := cli.Close(); err != nil {
			log.Err(err)
		}
	}()

	key, _, err := cli.Scan(ctx, 0, fmt.Sprintf("user:%v*|*", ident.GetUserIdentifier()), 1).Result()
	if err != nil {
		return nil, fmt.Errorf("not able to scan %w", err)
	}

	if len(key) == 0 {

		key, _, err = cli.Scan(ctx, 0, fmt.Sprintf("*|mail:%v", ident.GetUserIdentifier()), 1).Result()
		if err != nil {
			return nil, fmt.Errorf("not able to scan %w", err)
		}

		if len(key) == 0 {
			return nil, fmt.Errorf("user not found on redis database")
		}

	}

	return &key[0], nil
}

func getUserPg(ident models.UserIdentifier) (*models.User, error) {

	db := utils.PgConnect()

	query := fmt.Sprintf("SELECT USER_ID, USER_NAME, USER_PESS, USER_PASS, USER_MAIL FROM TB_USER WHERE UPPER(USER_NAME) = "+
		"UPPER(RTRIM('%[1]v')) OR USER_MAIL = UPPER(RTRIM('%[1]v'))", ident.GetUserIdentifier())

	rows, err := utils.PgQuery(db, 0, query)
	if err != nil {
		return nil, fmt.Errorf("error getting user data from database %w", err)
	}

	defer func() {
		err := rows.Close()
		log.Err(err)
	}()

	var id int
	var name, pess, pass, mail string

	if !rows.Next() {
		return nil, errors.New("value not found")
	}

	if err := rows.Scan(&id, &name, &pess, &pass, &mail); err != nil {
		return nil, fmt.Errorf("error scanning rows %w", err)
	}

	user := models.User{
		Id:       id,
		Username: name,
		PessName: pess,
		UserMail: mail,
		Password: pass,
	}

	return &user, nil
}

func createUserRedis(user models.User) error {

	ctx := context.Background()
	cli, err := utils.RedisCli()
	if err != nil {
		return fmt.Errorf("error creating redis client %w", err)
	}

	usr, _ := getUserRedis(user)
	if usr != nil {
		return fmt.Errorf("problem creating redis user %w", err)
	}

	hash, err := utils.CreateHash(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing user password %w", err)
	}

	key := fmt.Sprintf("user:%v|mail:%v", user.Username, user.UserMail)
	_, err = cli.RPush(ctx, key, user.Id, user.Username, user.PessName, user.UserMail, string(hash)).Result()
	if err != nil {
		return fmt.Errorf("error pushing into redis database %w", err)
	}

	return nil
}

func createUserPg(user models.User) (*int, error) {

	db := utils.PgConnect()

	rows, err := utils.PgQuery(db, 0, fmt.Sprintf("SELECT USER_ID FROM TB_USER WHERE USER_NAME = '%v' OR USER_MAIL = '%v';", user.Username, user.UserMail))
	if err != nil {
		return nil, fmt.Errorf("error searching for existing user on pg database %w", err)
	}

	defer func() {
		err := rows.Close()
		log.Err(err)
	}()

	if rows.Next() {
		return nil, fmt.Errorf("user already exists, not possible to create")
	}

	hash, err := utils.CreateHash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing user password %w", err)
	}

	insert := "INSERT INTO TB_USER (USER_NAME, USER_PASS, USER_PESS, USER_MAIL) " +
		fmt.Sprintf(" VALUES('%v', '%v', '%v', '%v');", user.Username, string(hash), user.PessName, user.UserMail)

	_, err = utils.PgQuery(db, 1, insert)
	if err != nil {
		return nil, fmt.Errorf("error inserting user on pg database")
	}

	usr, err := getUserPg(user)
	if err != nil {
		return nil, err
	}

	return &usr.Id, nil
}

func deleteUserRedis(ident models.UserIdentifier) error {

	key, err := scanUserRedis(ident)
	if err != nil {
		return err
	}

	cli, err := utils.RedisCli()
	if err != nil {
		return err
	}

	i := cli.Del(context.Background(), *key)
	if i == nil {
		return fmt.Errorf("error deleting redis user")
	}

	return nil
}

func deleteUserPg(ident models.UserIdentifier) error {

	db := utils.PgConnect()

	user, err := getUserPg(ident)
	if err != nil || user == nil {
		return fmt.Errorf("error deleting pg user %w", err)
	}

	_, err = utils.PgQuery(db, 3, fmt.Sprintf("DELETE FROM TB_USER WHERE USER_NAME = '%[1]v' OR USER_MAIL = '%[1]v'", ident.GetUserIdentifier()))
	if err != nil || user == nil {
		return fmt.Errorf("error deleting pg user %w", err)
	}

	return nil
}
