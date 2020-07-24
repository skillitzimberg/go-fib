import React from 'react';
import './FibInput.css';

export const FibInput = () => {
  console.log('FibInput rendered');
  const [fibNum, setFibNum] = React.useState(0);
  const [fibSeq, setFibSeq] = React.useState([]);

  function handleSubmit(event) {
    getFibSeq();
    var form = document.getElementById('fibForm');
    form.reset();
    event.preventDefault();
  }

  async function getFibSeq() {
    const fibSeq = await fetch(`/api/fibonacci/${fibNum}`).then(response =>
      response.json(),
    );
    setFibSeq(
      fibSeq
        ? fibSeq.join(', ')
        : 'There was a problem getting the sequence from the server. You might have requested too many values.',
    );
  }

  function onChange(event) {
    setFibNum(event.target.value);
  }

  return (
    <>
      <form onSubmit={handleSubmit} id="fibForm">
        <label htmlFor="fib">How many Fibonacci digits do you want?</label>
        <input type="number" min="1" id="fib" onChange={onChange} />
        <input type="submit" value="Enter" />
      </form>
      <p>Results:</p>
      <p id="results">{fibSeq}</p>
    </>
  );
};
