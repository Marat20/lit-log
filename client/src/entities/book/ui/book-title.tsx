import { useQuery } from '@tanstack/react-query';
import { FC, useEffect } from 'react';
import { useNavigate } from 'react-router';
import { fetchBook } from '../api/api';

export const BookTitle: FC = () => {
  const navigate = useNavigate();

  const { data } = useQuery({
    queryKey: ['books'],
    queryFn: fetchBook,
  });

  useEffect(() => {
    if (!data?.book) {
      navigate('/new');
    }
  });

  return (
    <section>
      <h1>{data?.book[data.book.length - 1].title}</h1>
      <h2>{data?.book[data.book.length - 1].author}</h2>
    </section>
  );
};
