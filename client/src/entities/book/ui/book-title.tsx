import { useQuery } from "@tanstack/react-query";
import { fetchBook } from "../model/api/api";

export const BookTitle = () => {
  const { data } = useQuery({
    queryKey: ["title"],
    queryFn: fetchBook,
  });
  
  return (
    <div>
      <h1>{data?.title}</h1>
      <h2>{data?.author}</h2>
    </div>
  );
};
