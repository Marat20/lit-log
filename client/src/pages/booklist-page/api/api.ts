import { Book } from "@/entities/book";
import { URL } from "@/shared/api/api";
import { userId } from "@/shared/consts/telegram";

export const fetchBooks = async (): Promise<{ books: Book[] } | undefined> => {
  try {
    const response = await fetch(`${URL}/${userId}/books`);
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};

export const setCurrentBook = async ({ bookId }: { bookId: string }) => {
  try {
    const response = await fetch(`${URL}/${userId}/setBook/${bookId}`, {
      method: "PUT",
    });
    const result = await response.json();

    return result;
  } catch (error) {
    console.error(error);
  }
};
