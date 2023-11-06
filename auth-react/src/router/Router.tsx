import { BrowserRouter, Route, Routes } from "react-router-dom";
import {Login,LoginPage} from "../components/Login";
import {Index} from "../components/Index";
import {Sandbox} from "../components/Sandbox"
import {Home} from "../home/Home";
import Page from "../components/Page";

export const Router = () => {
    return (
        <BrowserRouter>
        <div className="header">
            <LoginPage />
        </div>
        <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/login" element={<Login />} />
            <Route path="/home" element={<Home /> } />
            <Route path="/page" element={<Page />} />
            <Route path="/sandbox" element={<Sandbox />} />
        </Routes>
    </BrowserRouter>
    )
};