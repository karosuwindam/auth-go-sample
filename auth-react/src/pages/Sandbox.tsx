import axios from "axios";
import { Link } from "react-router-dom";
import "../public/css/Form.css";

import { useState, useContext, useEffect, FormEventHandler } from "react";
import { GetUser } from "../modules/UserEdit";

interface UserRerister {
  name: string;
  password: string;
  authority: number;
}

export const UserAdd: React.FC = () => {
  // const [user, setUser] = useState<UserRerister | null>(null);
  const datalist = {
    auth: [
      { id: 0, name: "guest" },
      { id: 1, name: "user" },
      { id: 2, name: "admin" },
    ],
  };

  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [authority, setAuthority] = useState(0);

  const data = JSON.parse(JSON.stringify(datalist));
  const mAuth = Object.keys(data.auth).map((key) => data.auth[key]);
  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };
  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setAuthority(Number(e.target.value));
  };
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const jsondata: UserRerister = {
      name: name,
      password: password,
      authority: authority,
    };
    console.log(JSON.stringify(jsondata));
  };

  return (
    <div className="container">
      <form onSubmit={handleSubmit}>
        <div>
          <div className="spacer" />
          <label>名前：</label>
          <input type="text" onChange={handleNameChange} name="" id="" />
        </div>
        <div>
          <div className="spacer" />
          <label>パスワード：</label>
          <input
            type="password"
            onChange={handlePasswordChange}
            name=""
            id=""
          />
        </div>
        <div>
          <div className="spacer" />
          <label>権限：</label>
          <select onChange={(e) => handleChange(e)}>
            {mAuth.map((auth) => (
              <option value={auth.id}>{auth.name}</option>
            ))}
          </select>
        </div>
        <div>
          <div className="spacer" />
          <input type="submit" value="Submit" />
        </div>
      </form>
    </div>
  );
};

export const Sandbox = () => {
  return (
    <div>
      <h1>Sandbox</h1>
      <UserAdd />
      <h2>view test module</h2>
      <Link to="/" className="App-link">
        back to index
      </Link>
    </div>
  );
};
