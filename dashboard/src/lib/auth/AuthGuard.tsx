import { Navigate } from "react-router-dom";
import { ReactNode, useContext } from "react";
import { authContext } from ".";

type Props = {
  link: string;
  children: ReactNode;
};

const Protected = ({ link = "/", children }: Props) => {
  const auth = useContext(authContext)

  if (auth?.user) {
    return <Navigate to={link} replace />;
  }
  return children;
};
export default Protected;
