import { Book } from "@/entities/book";
import { TrashIcon } from "@/shared/assets/icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FC } from "react";
import { deleteBook } from "../api/api";
import cls from "./delete-book.module.scss";

interface DeleteBookProps {
  bookId: string;
}

export const DeleteBook: FC<DeleteBookProps> = (props) => {
  const { bookId } = props;

  const queryClient = useQueryClient();

  const { mutate } = useMutation({
    mutationFn: deleteBook,
    onMutate: async (newData) => {
      await queryClient.cancelQueries({ queryKey: ["bookList"] });
      const previousData = queryClient.getQueryData<Book[]>(["bookList"]);

      queryClient.setQueryData(["bookList"], (old: { books: Book[] }) => ({
        ...old,
        books: old.books.filter((book) => book.id !== newData.bookId),
      }));

      return { previousData };
    },
  });

  const handleDeleteBook = () => {
    mutate({ bookId });
  };

  return (
    <button className={cls.btn_delete} onClick={handleDeleteBook}>
      <TrashIcon />
    </button>
  );
};
