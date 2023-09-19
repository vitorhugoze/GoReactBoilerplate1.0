const loginHandler = async function (user: string, pass: string): Promise<boolean> {
  const reqObj = {
    user: user,
    pass: pass,
  };

  const res = await fetch('/login', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(reqObj),
  });

  var isUserAuth: boolean;

  if (res.status === 200) {
    isUserAuth = true;
  } else {
    isUserAuth = false;
  }

  return new Promise((resolve) => {
    resolve(isUserAuth);
  });
};

export default loginHandler