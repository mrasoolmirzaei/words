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
    <>
      <div className="position-absolute top-50 start-50 translate-middle">
        <div className="d-flex justify-content-between mb-3 gap-2">
          <AddSynonym />
          <AddWord />
        </div>
        <SearchBar onSearch={handleSearch} />
        <SearchResults results={searchResults} />
        {loading && <Loading />}
      </div>
      <ToastContainer position="bottom-left" />
    </>
  );
};

export default App;
