import { BookListPage } from '@/pages/booklist-page';
import { NewBookPage } from '@/pages/newbook-page';
import { ProgressPage } from '@/pages/progress-page';
import { Footer } from '@/shared/ui/footer/footer';
import { Route, Routes } from 'react-router';

export const App = () => {
  return (
    <main className='main'>
      <Routes>
        <Route index element={<ProgressPage />} />
        <Route path='/books' element={<BookListPage />} />
        <Route path='/new' element={<NewBookPage />} />
      </Routes>
      <Footer />
    </main>
  );
};
