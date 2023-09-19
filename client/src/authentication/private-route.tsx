import { ReactElement, useEffect, useState, useContext } from 'react';
import { Outlet, useNavigate } from 'react-router-dom';
import { AuthContext } from './auth-context';

const PrivateRoute = (): ReactElement => {
  const [authenticated, setAuthenticated] = useState(false);
  const navigate = useNavigate();

  const context = useContext(AuthContext);

  const { SetAuthValue } = context!;

  useEffect(() => {
    fetch('/authuser').then((res) => {
      if (res.status !== 200) {
        SetAuthValue(false);
        navigate('/');
      } else {
        SetAuthValue(true);
        setAuthenticated(true);
      }
    });
  }, []);

  if (!authenticated) {
    return <h1 className="flex w-full items-center justify-center text-3xl text-slate-700">Loading......</h1>;
  }

  return <Outlet />;
};

export default PrivateRoute;
