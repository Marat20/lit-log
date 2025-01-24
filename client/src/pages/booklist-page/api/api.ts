import { Book } from "@/entities/book";
import { URL } from "@/shared/api/api";

export const fetchBooks = async (): Promise<{ books: Book[] } | undefined> => {
  try {
    const response = await fetch(`${URL}/books`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};

export const setCurrentBook = async ({ bookId }: { bookId: string }) => {
  try {
    const response = await fetch(`${URL}/setBook/${bookId}`, {
      method: "PUT",
    });
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
