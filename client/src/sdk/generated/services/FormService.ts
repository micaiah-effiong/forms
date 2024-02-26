/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Form } from '../models/Form';
import type { FormData } from '../models/FormData';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class FormService {
  /**
   * Get all forms
   * @returns Form Successfully get all forms
   * @throws ApiError
   */
  public static getForms(): CancelablePromise<Array<Form>> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/forms',
    });
  }
  /**
   * Create a new form
   * @param requestBody
   * @returns Form Successfully create a new form
   * @throws ApiError
   */
  public static postForms(
    requestBody?: {
      name?: string;
    },
  ): CancelablePromise<Array<Form>> {
    return __request(OpenAPI, {
      method: 'POST',
      url: '/forms',
      body: requestBody,
      mediaType: 'application/json',
    });
  }
  /**
   * Get a form
   * @param id
   * @returns Form Successfully get a form
   * @throws ApiError
   */
  public static getFormById(
    id: string,
  ): CancelablePromise<Form> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/forms/{id}',
      path: {
        'id': id,
      },
    });
  }
  /**
   * Delete a form
   * @param id
   * @returns Form Successfully delete a form
   * @throws ApiError
   */
  public static deleteForms(
    id: string,
  ): CancelablePromise<Form> {
    return __request(OpenAPI, {
      method: 'DELETE',
      url: '/forms/{id}',
      path: {
        'id': id,
      },
    });
  }
  /**
   * Get a form data
   * @param id
   * @returns FormData Successfully get a form data
   * @throws ApiError
   */
  public static getFormsData(
    id: string,
  ): CancelablePromise<Array<FormData>> {
    return __request(OpenAPI, {
      method: 'GET',
      url: '/forms/{id}/data',
      path: {
        'id': id,
      },
    });
  }
  /**
   * Add a form data
   * @param id
   * @param requestBody
   * @returns FormData Successfully add a form data
   * @throws ApiError
   */
  public static postFormsData(
    id: string,
    requestBody?: Record<string, any>,
  ): CancelablePromise<FormData> {
    return __request(OpenAPI, {
      method: 'POST',
      url: '/forms/{id}/data',
      path: {
        'id': id,
      },
      body: requestBody,
      mediaType: 'application/json',
    });
  }
}
