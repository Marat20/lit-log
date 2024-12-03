import { useQuery } from "@tanstack/react-query";
import { fetchBook } from "../model/api/api";

export const BookProgress = () => {
  const { data: book } = useQuery({
    queryKey: ["title"],
    queryFn: fetchBook,
  });

  return <progress value={book?.currentPage} max={book?.totalPages} />;
};
