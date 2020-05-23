import React, { useState, createContext, useCallback, useContext } from 'react';
import ChatContext from './chat';

export interface IUser {
  name: string;
}

export interface IUserContext {
  user: IUser;
  join: (username: string) => void;
}

export interface UserContextProps {
  children: JSX.Element;
}

const UserContext = createContext<IUserContext>(null);

UserContext.displayName = 'UserContext';

export function UserContextProvider(props: UserContextProps): JSX.Element {
  const [user, setUser] = useState<IUser>(null);
  const { connect } = useContext(ChatContext);

  const join = useCallback((username: string) => {
    setUser({
      name: username
    });
    connect();
  }, []);

  const value: IUserContext = {
    user,
    join
  }

  return (
    <UserContext.Provider value={value}>
      {props.children}
    </UserContext.Provider>
  );
}

export default UserContext;
