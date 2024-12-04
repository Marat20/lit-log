import { BookProgress, BookTitle } from "@/entities/book";
import { UpdatePagesRead } from "@/features/update-pages-read";
import { FC } from "react";

export const ProgressPage: FC = () => {
  return (
    <>
      <BookTitle />
      <BookProgress />
      <UpdatePagesRead />
    </>
  );
};
