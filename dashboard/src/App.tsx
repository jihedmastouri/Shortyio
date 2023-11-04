import { lazy, useContext } from "react";
import { Routes, Route } from "react-router-dom";
import axios from "axios";

import useAuth from "@hooks/auth";
import { AuthContext } from "@lib/auth";

import NotFound from "pages/notFound";
import Home from "pages/home";
import Welcome from "pages/Welcome";
import Posts from "pages/home/Posts";
import Users from "pages/home/Users";
import Media from "pages/home/Media";
import HomePage from "pages/home/HomePage";
import Interactions from "pages/home/Interactions";
import EditPost from "pages/home/Posts/Edit";
import Login from "pages/login";


const NewDomain = lazy(() => import("pages/NewDomain"));

function App() {
  const { refreshToken } = useAuth();
  const { accessToken } = useContext(AuthContext);

  axios.defaults.baseURL = "http://localhost:42069/";

  axios.interceptors.request.use(
    (config) => {
      if (accessToken)
        axios.defaults.headers.common[
          "Authorization"
        ] = `Bearer ${accessToken}`;
      return config;
    },
    (error) => Promise.reject(error)
  );

  axios.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config;

      if (
        error.response &&
        error.response.status === 401 &&
        !originalRequest._retry
      ) {
        // Prevent infinite loop
        originalRequest._retry = true;

        try {
          await refreshToken();
          return axios(originalRequest);
        } catch (refreshError) {
          console.error(refreshError);
          return Promise.reject(refreshError);
        }
      }
      return Promise.reject(error);
    }
  );

  return (
    <Routes>
      <Route path="" element={<Home />}>
        <Route index element={<HomePage />} />
        <Route path="/posts" element={<Posts />} />
        <Route path="/posts/:id" element={<EditPost />} />
        <Route path="/users" element={<Users />} />
        <Route path="/media" element={<Media />} />
        <Route path="/interactions" element={<Interactions />} />
      </Route>
      <Route path="/welcome" element={<Welcome />} />
      <Route path="/login" element={<Login />} />
      <Route path="/new-domain" element={<NewDomain />} />
      <Route path="/404" element={<NotFound />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}

export default App;
