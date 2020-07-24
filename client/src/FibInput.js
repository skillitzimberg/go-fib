import React from 'react';

export const FibInput = () => {
  return (
    <>
      <form id="fibForm">
        <label htmlFor="fib">How many Fibonacci digits do you want?</label>
        <input type="number" min="1" id="fib" />
        <input type="submit" value="Enter" />
      </form>
      <p>Results</p>
    </>
  );
};
