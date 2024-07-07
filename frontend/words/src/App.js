import ActionButtons from "./components/actionButtons/ActionButtons";
import SearchBar from "./components/searchBar/SearchBar";
import useSearchWord from "./services/hook/useSearchWord";

const App = () => {

  const { searchResults } = useSearchWord();

  return (
    <div className="position-absolute top-50 start-50 translate-middle">
      <SearchBar />
      <ActionButtons />
      {searchResults && <div>Search Results: {JSON.stringify(searchResults)}</div>}
    </div>
  );
};

export default App;
