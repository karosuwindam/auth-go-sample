import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Login, LoginPage } from "../modules/Login";
import { Index } from "../pages/Index";
import { Home } from "../pages/Home";
import { Page } from "../pages/Page";
import { UserList, Edit, UserView } from "../pages/Edit";
import { Sandbox } from "../pages/Sandbox";

export const Router = () => {
  return (
    <BrowserRouter>
      <div className="header">
        <LoginPage />
      </div>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/login" element={<Login />} />
        <Route path="/edit" element={<Edit />} />
        <Route path="/user/List" element={<UserList />} />
        <Route path="/user/:id" element={<UserView />} />
        <Route path="/home" element={<Home />} />
        <Route path="/page" element={<Page />} />
        <Route path="/sandbox" element={<Sandbox />} />
      </Routes>
    </BrowserRouter>
  );
};
