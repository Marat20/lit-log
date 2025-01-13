import { URL } from '@/shared/api/api';
import { userId } from '@/shared/consts/consts';
import { Book } from '../model/types/book';

export const fetchBook = async (): Promise<
  { book: Book[]; pagesReadToday: number } | undefined
> => {
  try {
    const response = await fetch(`${URL}/${userId}/init}`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
