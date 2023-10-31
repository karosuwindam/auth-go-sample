import { BrowserRouter, Route, Routes } from "react-router-dom";
import {Login} from "../components/Login";
import {Index} from "../components/Index";
import {Sandbox} from "../components/Sandbox"
import Page from "../components/Page";

export const Router = () => {
    return (
        <BrowserRouter>
        <Routes>
            <Route path="/" element={<Index />} />
            <Route path="/login" element={<Login />} />
            <Route path="/page" element={<Page />} />
            <Route path="/sandbox" element={<Sandbox />} />
        </Routes>
    </BrowserRouter>
    )
};