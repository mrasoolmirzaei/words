const Loading = () => (
  <div
    className="position-fixed top-0 start-0 w-100 h-100 
                d-flex justify-content-center align-items-center bg-light 
                bg-opacity-75"
  >
    <div className="spinner-border" role="status">
      <span className="visually-hidden">Loading...</span>
    </div>
  </div>
);

export default Loading;
