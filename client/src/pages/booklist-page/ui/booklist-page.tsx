import { BookCard } from "@/entities/book";
import { useQuery } from "@tanstack/react-query";
import { fetchBooks } from "../api/api";

const BookListPage = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["bookList"],
    queryFn: fetchBooks,
  });

  // const queryClient = useQueryClient();

  // const { mutate } = useMutation({
  //   mutationFn: setCurrentBook,
  //   onSuccess: (data: ReturnData) => {
  //     queryClient.setQueryData(["book"], {
  //       currentBook: data.currentBook,
  //       pagesReadToday: data.pagesReadToday,
  //     });
  //   },
  // });

  // const onSubmit = (bookId: string) => {
  //   mutate({ bookId });
  // };

  if (data?.books.length == 0) {
    return <div>No books</div>;
  }

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
