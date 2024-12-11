import { useQuery } from "@tanstack/react-query";
import { FC } from "react";
import { fetchBook } from "../api/api";

export const BookTitle: FC = () => {
  const { data } = useQuery({
    queryKey: ["book"],
    queryFn: fetchBook,
  });

  return (
    <section>
      <h1>{data?.book.title}</h1>
      <h2>{data?.book.author}</h2>
    </section>
  );
};
