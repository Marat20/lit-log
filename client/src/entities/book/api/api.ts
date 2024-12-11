import { URL } from '@/shared/api/api';
import { Book } from '../model/types/book';

export const fetchBook = async (): Promise<
  { book: Book; pagesReadToday: number } | undefined
> => {
  try {
    const response = await fetch(`${URL}/YJWNnpoLOSBq6HVcK7oW9`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
