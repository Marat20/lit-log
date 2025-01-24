import { useMutation } from "@tanstack/react-query";
import { FC } from "react";
import { deleteBook } from "../api/api";

interface DeleteBookProps {
  bookId: string;
  className: string
}

export const DeleteBook: FC<DeleteBookProps> = (props) => {
  const { bookId, className } = props;

  const { mutate } = useMutation({
    mutationFn: deleteBook,
  });

  const handleDeleteBook = () => {
    mutate({ bookId });
  };

  return <button className={className} onClick={handleDeleteBook}>Del</button>;
};
