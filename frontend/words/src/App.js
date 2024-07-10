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
    <div className="m-5">
      <h1 className="text-center p-5">Welcome to Words</h1>
      <div className="position-absolute top-50 start-50 translate-middle">
        <AddWord />
        <AddSynonym />
        <SearchBar onSearch={handleSearch} />
        <SearchResults results={searchResults} />
      </div>
      {loading && <Loading />}
      <ToastContainer position="bottom-left" />
    </div>
  );
};

export default App;
