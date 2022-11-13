import React from 'react';
import BasicTable from './components/BasicTable';
import { QueryClient, QueryClientProvider } from "react-query";

function App() {
  const queryClient = new QueryClient();

  return (
    <div className="App">
      <QueryClientProvider client={queryClient}>
        <BasicTable />
      </QueryClientProvider>
    </div>
  );
}

export default App;
