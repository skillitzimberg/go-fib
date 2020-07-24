import React from 'react';

export const FibInput = () => {
  const [fibNum, setFibNum] = React.useState(0);
  const [fibSeq, setFibSeq] = React.useState([]);

  function handleSubmit(event) {
    console.log('submitted');
    var form = document.getElementById('fibForm');
    form.reset();
    event.preventDefault();
  }

  function onChange(event) {
    setFibNum(event.target.value);
  }

  return (
    <>
      <form id="fibForm" onSubmit={handleSubmit}>
        <label htmlFor="fib">How many Fibonacci digits do you want?</label>
        <input type="number" min="1" id="fib" onChange={onChange} />
        <input type="submit" value="Enter" />
      </form>
      <p>Results</p>
    </>
  );
};
