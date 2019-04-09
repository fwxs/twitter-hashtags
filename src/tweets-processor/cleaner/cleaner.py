from nltk.corpus import stopwords
import nltk
import re


class Broom:
    def __init__(self, language):
        self._rt_regex = re.compile('RT\s@\w{1,}:\s')
        self._hashtag_regex = re.compile('\s{1,}#\w{1,}\s{1,}')
        self._no_alpha_regex = re.compile(r'[!\"#$¿¡*%&\'\(\)*+,-./:;<=>?@\[\\\]^_`\{|\}~]')
        self.stopwords = set()
        self.initializer(language)

    def clean_rt(self, line):
        return re.sub(self._rt_regex, ' ', line)

    def clean_hashtag(self, line):
        return re.sub(self._hashtag_regex, ' ', line)

    def clean_no_alpha(self, line):
        return re.sub(self._no_alpha_regex, ' ', line)

    def clean_stopwords(self, line):
        return " ".join((word for word in line.split() if word not in self.stopwords))

    def cleaner(self, line):
        clean_data = self.clean_no_alpha(self.clean_hashtag(self.clean_rt(line)))
        return self.clean_stopwords(clean_data) if len(self.stopwords) == 0 else clean_data

    def initializer(self, language):
        if nltk.download('stopwords'):
            self.stopwords.update(stopwords.words(language))

