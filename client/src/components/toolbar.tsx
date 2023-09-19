import { useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaReact } from 'react-icons/fa';
import loginHandler from '../handlers/login-handler';
import { AuthContext } from '../authentication/auth-context';

function LoginComponent() {
  const [clickedLogin, setClickedLogin] = useState(false);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const navigate = useNavigate();

  return (
    <div
      onClick={(e) => {
        if (e.target === e.currentTarget) {
          setClickedLogin(false);
        }
      }}
      className="flex w-full justify-end"
    >
      <div className={`mr-24 flex items-center`}>
        <input onChange={(e) => setUsername(e.target.value)} type="text" name="" id="user-input" placeholder="Username" className={`${clickedLogin ? '' : 'cursor-default opacity-0 '} w-45 mr-3 flex rounded border bg-zinc-300 align-middle transition-all duration-300 focus:outline-none focus:ring-0`} />
        <input onChange={(e) => setPassword(e.target.value)} type="password" name="" id="password-input" placeholder="Password" className={`${clickedLogin ? '' : 'cursor-default opacity-0 '} w-45 flex rounded border bg-zinc-300 align-middle transition-all duration-300 focus:outline-none focus:ring-0`} />
      </div>
      <button
        onClick={() => {
          if (clickedLogin) {
            loginHandler(username, password).then((auth) => {
              if (auth) {
                navigate('/userpage');
              } else {
                navigate('/');
              }
            });
          } else {
            setClickedLogin(true);
          }
        }}
        className={`${clickedLogin ? 'translate-x-24 transform ' : ''} absolute right-28 top-2 h-[28px] w-[61px] select-none justify-center rounded-md bg-cyan-700 align-middle text-stone-200 transition-all duration-500 hover:bg-cyan-800`}
      >
        Login
      </button>
      <button
        onClick={() => {
          navigate('/signup');
        }}
        className={`${clickedLogin ? 'hidden' : ''} absolute right-3 top-2 flex h-[28px] w-[81px] select-none justify-center rounded-md border-2 border-stone-200 align-middle text-stone-200 hover:bg-zinc-900`}
      >
        Sign up
      </button>
    </div>
  );
}

function LoggedComponent() {
  const [clickedLogo, setClickedLogo] = useState(false);

  useEffect(() => {
    var latMen = document.getElementById('lateral-menu');

    if (latMen != null) {
      if (clickedLogo) {
        latMen.style.width = '12rem';
      } else {
        latMen.style.width = '0px';
      }
    }
  }, [clickedLogo]);

  return (
    <div className="flex items-center">
      <FaReact className={`ml-3 h-[32px] w-[32px] cursor-pointer fill-cyan-400 duration-300 ${clickedLogo ? 'rotate-180' : ''} `} onClick={() => setClickedLogo(!clickedLogo)} />
    </div>
  );
}

function Toolbar() {
  const context = useContext(AuthContext)!;
  const { isAuth } = context;

  return (
    <div className="flex h-[45px] w-full bg-zinc-800">
      {isAuth ? <LoggedComponent /> : null}
      {!isAuth ? <LoginComponent /> : null}
    </div>
  );
}

export default Toolbar;
