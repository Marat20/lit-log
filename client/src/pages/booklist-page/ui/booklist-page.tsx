import { BookCard } from '@/entities/book';
import { useQuery } from '@tanstack/react-query';
import { fetchBooks } from '../api/api';

// TODO

const BookListPage = () => {
  const { data } = useQuery({
    queryKey: ['books'],
    queryFn: fetchBooks,
  });

  return (
    <>
      {data?.books.map((item) => (
        <BookCard key={item.ID} bookData={item} />
      ))}
    </>
  );
};

export default BookListPage;
