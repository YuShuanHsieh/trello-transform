import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import { ObjectView } from '../../shared/objectView';
import style from './result.module.css';

function ResultComponent({ result }) {
  if (Object.keys(result).length === 0) {
    return (
      <div className={style.preparing}>
        <p>Please upload your file first</p>
      </div>
    );
  }
  return (
    <ObjectView object={result} property="Result" />
  );
}

ResultComponent.propTypes = {
  result: PropTypes.objectOf(PropTypes.object).isRequired,
};

function mapStateTpProps(state) {
  return {
    result: state,
  };
}

export const Result = connect(mapStateTpProps)(ResultComponent);
