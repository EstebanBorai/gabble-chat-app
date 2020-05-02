import React, { useContext } from 'react';
import './login.scss';
import ChatContext, { ChatContextInterface } from '../../contexts/chat';

const Login = (): JSX.Element => {
  const [username, setUsername] = React.useState('');

  const { join } = useContext<ChatContextInterface>(ChatContext);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();

    join(username);
  }

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setUsername(event.target.value);
  }

  return (
    <section className="application-section" id="login">
      <article>
        <h3>Join the session!</h3>
        <form action="" onSubmit={handleSubmit}>
          <input type="text" name="username" value={username} onChange={handleChange} />
          <label htmlFor="join">
            <button type="submit" className="primary">Join</button>
          </label>
        </form>
      </article>
    </section>
  );
}

export default Login;
