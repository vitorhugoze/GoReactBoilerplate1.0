import { ReactElement, useContext, useState, useEffect } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import { AuthContext } from './auth-context';

const RedirectUserRoute = (): ReactElement => {
  const [authenticated, setAuthenticated] = useState(false);

  const context = useContext(AuthContext);
  const navigate = useNavigate();

  const { SetAuthValue } = context!;

  useEffect(() => {
    fetch('/authuser').then((res) => {
      if (res.status !== 200) {
        SetAuthValue(false);
        setAuthenticated(true);
      } else {
        SetAuthValue(true);
        setAuthenticated(true);
        navigate('/userpage');
      }
    });
  }, []);

  if (!authenticated) {
    return <h1 className="flex w-full items-center justify-center text-3xl text-slate-700">Loading......</h1>;
  }

  return <Outlet />;
};

export default RedirectUserRoute;
