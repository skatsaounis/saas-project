export interface Question {
  id: string;
  question: string;
  answer: string;
}

export interface QResponse {
  status: number;
  message: string;
  data: QuestionData;
}

export interface QuestionData {
  data: Question[];
}
