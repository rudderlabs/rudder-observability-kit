import { Label } from '../labels';

export interface Logger {
  debug(message: string, ...labels: Label[]): void;
  info(message: string, ...labels: Label[]): void;
  warn(message: string, ...labels: Label[]): void;
  error(message: string, ...labels: Label[]): void;
  fatal(message: string, ...labels: Label[]): void;
  getLabels(): Label[];
  createLogger(...labels: Label[]): Logger;
}
