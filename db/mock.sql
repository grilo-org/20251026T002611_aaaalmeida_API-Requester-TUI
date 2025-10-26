INSERT INTO collection (name, description) VALUES
('local simples', 'description 1'),
('local complexas', 'description 2'),
('exterior', 'description 3'),
('faculdade', 'desc faculdade'),
('trabalho', 'desc trabalho'),
('pessoal', 'desc pessoal');

INSERT INTO request (name,                  url,                                          method_id,  collection_id) VALUES
                    ('teste GET',         'http://localhost:8080/',                       1,          1),
                    ('teste POST',        'http://localhost:8080/',                       2,          1),
                    ('teste PUT',         'http://localhost:8080/',                       3,          1),
                    ('teste DELETE',      'http://localhost:8080/',                       4,          1),
                    ('teste PATCH',       'http://localhost:8080/',                       5,          1),
                    ('teste HEAD',        'http://localhost:8080/',                       6,          1),
                    ('teste TRACE',       'http://localhost:8080/',                       7,          1),
                    ('teste OPTIONS',     'http://localhost:8080/',                       8,          1),
                    ('ola mundo',         'http://localhost:8080/',                       1,          1),
                    ('date_time',         'http://localhost:8080/time',                   1,          1),
                    ('teste params',      'http://localhost:8080/id=abcd',                1,          2),
                    ('example',           'https://www.example.com',                      1,          3),
                    ('json placeholder',  'https://jsonplaceholder.typicode.com/posts/1', 1,          3),
                    ('json placeholder',  'https://jsonplaceholder.typicode.com/posts',   2,          3),
                    ('json placeholder',  'https://jsonplaceholder.typicode.com/posts',   3,          3);

INSERT INTO request (name,                            url,                               method_id,  collection_id,  body,                                                                           body_type,  headers) VALUES
                    ('teste retorna body',            'http://localhost:8080/body',      2,          2,              '{"valorBody": "abc", "num": 123}',                                             2,          null),
                    ('teste retorna body',            'http://localhost:8080/headers',   2,          2,              null,                                                                           0,          '{"Hello": "world"}'),
                    ('teste retorna parte de body',   'http://localhost:8080/obj',       2,          2,              '{"ignora": true, "valor": [{"num": 123}, "resto": true}, {"alo": "mundo"}]',   2,          null),
                    ('teste retorn header e body',    'http://localhost:8080/all',       2,          2,              '{"hello": "world"}',                                                           2,          '{"headerValue": 123}');
