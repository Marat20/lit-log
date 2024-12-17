import { Link } from "react-router";
import cls from "./footer.module.scss";

export const Footer = () => {
  return (
    <footer className={cls.footer}>
      <Link to={"/books"}>Books</Link>
      <Link to={"/"}>Home</Link>
    </footer>
  );
};
