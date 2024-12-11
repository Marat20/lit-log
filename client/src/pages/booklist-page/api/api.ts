import { Book } from '@/entities/book';
import { URL } from '@/shared/api/api';

export const fetchBooks = async (): Promise<{ books: Book[] } | undefined> => {
  try {
    const response = await fetch(`${URL}`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
