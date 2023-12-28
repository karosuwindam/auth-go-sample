import axios from "axios";
import { Link } from "react-router-dom";

import { useState, useContext, useEffect } from "react";
import { GetUser } from "../modules/UserEdit";

export const Sandbox = () => {
  const json = GetUser(1);
  console.log(json);

  return (
    <div>
      <h1>Sandbox</h1>
      <h2>view test module</h2>
      <Link to="/" className="App-link">
        back to index
      </Link>
    </div>
  );
};
