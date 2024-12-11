import { BookListPageLazy } from '@/pages/booklist-page';
import { ProgressPageLazy } from '@/pages/progress-page';
import { Route, Routes } from 'react-router';

export const App = () => {
  return (
    <main className='main'>
      <Routes>
        <Route index element={<ProgressPageLazy />} />
        <Route path='/books' element={<BookListPageLazy />} />
      </Routes>
    </main>
  );
};
