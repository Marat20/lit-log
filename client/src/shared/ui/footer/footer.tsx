import { memo } from "react";
import { Link } from "react-router";
import cls from "./footer.module.scss";

export const Footer = memo(() => {
  return (
    <footer className={cls.footer}>
      <Link to={"/books"}>Books</Link>
      <Link to={"/new"}>New</Link>
      <Link to={"/"}>Home</Link>
    </footer>
  );
});
