import axios from "axios";
import { Link } from "react-router-dom";

import { useState, useContext, useEffect } from "react";
import { GetUser } from "../modules/UserEdit";

interface UserRerister {
  name: string;
  password: string;
  authority: number;
}

export const UserAdd: React.FC = () => {
  const [user, setUser] = useState <UserRerister | null>(null);
  return (
    <div>
      <div>user add</div>
      <div>add buttoon</div>
    </div>
  )
}

export const Sandbox = () => {
  const json = GetUser(1);
  console.log(json);

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
