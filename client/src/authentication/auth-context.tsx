import { createContext, ReactElement, useState } from 'react';

type AuthContent = {
  isAuth: boolean;
  SetAuthValue(val: boolean): void;
};

type ProviderContent = {
  children: ReactElement;
};

export const AuthContext = createContext<AuthContent | undefined>(undefined);

export function AuthProvider(content: ProviderContent): ReactElement {
  const [isAuth, setUserAuth] = useState<boolean>(false);

  function SetAuthValue(val: boolean) {
    setUserAuth(val);
  }

  return <AuthContext.Provider value={{ isAuth, SetAuthValue }}>{content.children}</AuthContext.Provider>;
}
