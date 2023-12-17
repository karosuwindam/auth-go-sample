import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Login, LoginPage } from "../modules/Login";
import { Index } from "../pages/Index";
import { Home } from "../pages/Home";
import { Page } from "../pages/Page";
import { Edit} from "../pages/Edit";

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
        <Route path="/home" element={<Home />} />
        <Route path="/page" element={<Page />} />
      </Routes>
    </BrowserRouter>
  );
};
