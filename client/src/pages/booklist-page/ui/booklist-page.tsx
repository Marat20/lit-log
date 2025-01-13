import { BookCard } from "@/entities/book";
import { useQuery } from "@tanstack/react-query";
import { fetchBooks } from "../api/api";

const BookListPage = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["bookList"],
    queryFn: fetchBooks,
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <>
      {data?.books.map((item) => (
        <BookCard key={item.id} bookData={item} />
      ))}
    </>
  );
};

export default BookListPage;
