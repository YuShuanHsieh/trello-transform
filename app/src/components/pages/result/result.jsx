import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

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

function mapStateTpProps(state) {
  return {
    uploaded: state.uploaded,
    result: state.result,
  };
}

export const Result = connect(mapStateTpProps)(ResultComponent);
