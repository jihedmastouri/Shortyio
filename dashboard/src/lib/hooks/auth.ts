import { useContext } from "react";
import { AuthContext } from "lib/auth";
import axios from "axios";

const useAuth = () => {

  const { accessToken: access_token, setAccessToken } = useContext(AuthContext);

  const login = async (email:string, password: string) => {
    axios.post('/api/login', { email, password }).then((res) => {
      setAccessToken(res.data.access_token);
    });
  };

  const logout = () => {
    axios.post('/api/logout');
    setAccessToken(null);
  };

  const isLoggedIn = () => {
    return !!access_token;
  };

  const refreshToken = async () => {
    return axios.post('/api/refresh').then((res) => {
      setAccessToken(res.data.access_token);
    });
  };

	return { login, logout, isLoggedIn, refreshToken } as const;
};

export default useAuth;
