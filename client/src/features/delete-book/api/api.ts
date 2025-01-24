import { URL } from "@/shared/api/api";

export const deleteBook = async ({ bookId }: { bookId: string }) => {
  try {
    await fetch(`${URL}/${bookId}`, {
      method: "DELETE",
    });
  } catch (error) {
    console.error(error);
  }
};
