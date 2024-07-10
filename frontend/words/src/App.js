import { ToastContainer } from "react-toastify";
import AddSynonym from "./components/actionButtons/AddSynonym";
import AddWord from "./components/actionButtons/AddWord";
import SearchBar from "./components/searchBar/SearchBar";
import SearchResults from "./components/searchResults/SearchResults";
import useSearchWord from "./services/hook/useSearchWord";
import Loading from "./components/loading/Loading";

const App = () => {
  const { loading, searchResults, handleSearch } = useSearchWord();

  return (
    <div className="d-flex flex-column align-items-center vh-100">
      <h1 className="text-center p-5">Welcome to Words</h1>
      <div className="d-flex flex-column align-items-center w-100 max-w-28rem mt-7">
        <AddWord />
        <AddSynonym />
        <SearchBar onSearch={handleSearch} />
        <div className="w-100 mt-3 overflow-auto max-h-50vh">
          <SearchResults results={searchResults} />
        </div>
      </div>
      {loading && <Loading />}
      <ToastContainer position="bottom-left" />
    </div>
  );
};

export default App;
