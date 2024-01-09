import axios from "axios";
import { Link } from "react-router-dom";
import "../public/css/Form.css";

import { useState, useContext, useEffect, FormEventHandler } from "react";
import { GetUser } from "../modules/UserEdit";

import { UserAdd } from "../modules/UserAdd";

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
