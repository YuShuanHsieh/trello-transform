import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createSelector } from 'reselect';

import { ObjectView } from '../../shared/objectView';
import style from './result.module.css';

function ResultComponent({ uploaded, result }) {
  if (!uploaded && Object.keys(result).length === 0) {
    return (
      <div className={style.preparing}>
        <span className={style.border}>Please upload your file first</span>
      </div>
    );
  }
  return (
    <div className={style.resultContainer}>
      <div className={style.result}>
        <h2>Result</h2>
        <ObjectView object={result} property="Result" />
      </div>
    </div>
  );
}

ResultComponent.propTypes = {
  uploaded: PropTypes.bool.isRequired,
  result: PropTypes.oneOfType([PropTypes.object, PropTypes.array]).isRequired,
};

const resultSelector = createSelector(
  state => state.result,
  (result) => {
    if (result.list && result.list.length > 0) {
      result.list.sort((a, b) => {
        const aDate = new Date(a.date);
        const bDate = new Date(b.date);
        return aDate.getTime() - bDate.getTime();
      });
      const list = result.list.map(item => (`${item.date} - ${item.title}`));
      return { ...result, list };
    }
    return result;
  },
);

function mapStateTpProps(state) {
  return {
    uploaded: state.uploaded,
    result: resultSelector(state),
  };
}

export const Result = connect(mapStateTpProps)(ResultComponent);
