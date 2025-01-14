import clsx from "clsx";
import { FC, memo, ReactNode } from "react";
import cls from "./error-message.module.scss";

interface ErrorMessageProps {
  children: ReactNode;
  className?: string;
}

export const ErrorMessage: FC<ErrorMessageProps> = memo((props) => {
  const { children, className } = props;

  return <p className={clsx(cls.error, className)}>{children}</p>;
});
