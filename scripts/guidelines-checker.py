#!/usr/bin/python3
import os
import re
import sys

_VERSION = (1, 0, 0)
_VERSION_STR = '.'.join([str(i) for i in _VERSION.__iter__()])

if __package__:
    CURRENT_PATH = os.path.abspath(os.path.dirname(__file__)).replace('/' + str(__package__), '', 1)
else:
    CURRENT_PATH = os.path.abspath(os.path.dirname(__file__))

if not CURRENT_PATH[-1]=='/':
    CURRENT_PATH += '/'

if re.search("/(scripts)/", CURRENT_PATH):
    ROOT_DIR = CURRENT_PATH.replace('scripts/', '')
else:
    ROOT_DIR = CURRENT_PATH.replace('bin/', '')

_GUIDELINES = {
    "github-templates": False,
    "github-actions": False,
    "github-docs": {
        "result": False,
        "files": [
            {
                "README.md": False
            },
            {
                "CHANGELOG.md": False
            },
            {
                "CODE_OF_CONDUCT.md": False
            },
            {
                "CONTRIBUTING.md": False
            },
            {
                "LICENSE.md": False
            }
        ]
    },
    "sonar": False,
    "sonar-action": False,
    "openapi": False,
    "restful": False,
    "docker": False,
    "editorconfig": False,
    "serverless": False,
    "logs-monitoring": False,
    "scripts": False,
    "environment-scripts": False,
    "terraform": False,
    "tests":
        {
            "result": False,
            "contexts": [
                {
                    "unit": False
                },
                {
                    "integration": False
                },
                {
                    "component": False
                },
            ]
        },
    "aws-ci-cd": False,
    "meta-files": [
        {
            ".projectrc": False
        },
        {
            ".dockerignore": False
        },
        {
            ".gitignore": False
        },
        {
            "docker-compose.yml": False
        }
    ]
}
_GD_KEYS = _GUIDELINES.keys()
_FINAL_RESULT = False


def repeat(word, m, n):
    # if number of characters greater than length of word.
    # set number of characters = length of word
    if m > len(word):
        m = len(word)

    repeat_word = word[:m]
    return repeat_word * n


class GuidelinesChecker:
    GUIDES_CONFIG = {
        "github-templates": {"callback": "has_github_templates", "required": True},
        "github-actions": {"callback": "has_github_actions", "required": True},
        "github-docs": {"callback": "has_github_docs", "required": True},
        "sonar": {"callback": "has_sonar", "required": True},
        "sonar-action": {"callback": "has_sonar_action", "required": True},
        "openapi": {"callback": "has_openapi", "required": False},
        "restful": {"callback": "has_restful", "required": False},
        "docker": {"callback": "has_docker", "required": True},
        "editorconfig": {"callback": "has_editorconfig", "required": True},
        "serverless": {"callback": "has_serverless", "required": False},
        "logs-monitoring": {"callback": "has_logs_monitoring", "required": False},
        "scripts": {"callback": "has_scripts", "required": True},
        "environment-scripts": {"callback": "has_environment_scripts", "required": True},
        "terraform": {"callback": "has_terraform", "required": False},
        "tests": {"callback": "has_tests", "required": True},
        "aws-ci-cd": {"callback": "has_aws_ci_cd", "required": False},
        "meta-files": {"callback": "has_meta_files", "required": False},
    }

    def check(self):
        global _GUIDELINES
        global _FINAL_RESULT
        for guide_key, object in self.GUIDES_CONFIG.items():
            callback = object['callback']
            if hasattr(self, str(callback)):
                result = getattr(self, callback)()
                if object["required"]:
                    if isinstance(result, bool):
                        if _FINAL_RESULT:
                            _FINAL_RESULT = True if result else False
                    elif isinstance(result, dict):
                        res = result['result']
                        if _FINAL_RESULT:
                            _FINAL_RESULT = True if res else False
                _GUIDELINES[guide_key] = result

    @staticmethod
    def has_github_templates():
        github_dir = os.path.isdir(ROOT_DIR + '.github/')
        github_pr_template = os.path.isfile(ROOT_DIR + '.github/pull_request_template.md')
        github_issue_template_dir = os.path.isdir(ROOT_DIR + '.github/ISSUE_TEMPLATE')
        return github_dir and github_pr_template and github_issue_template_dir

    @staticmethod
    def has_github_actions():
        github_dir = os.path.isdir(ROOT_DIR + '.github/')
        github_workflows_dir = os.path.isdir(ROOT_DIR + '.github/workflows')
        return github_dir and github_workflows_dir

    @staticmethod
    def has_github_docs():
        README = os.path.isfile(ROOT_DIR + 'README.md')
        CHANGELOG = os.path.isfile(ROOT_DIR + 'CHANGELOG.md')
        CODE_OF_CONDUCT = os.path.isfile(ROOT_DIR + 'CODE_OF_CONDUCT.md')
        CONTRIBUTING = os.path.isfile(ROOT_DIR + 'CONTRIBUTING.md')
        LICENSE = os.path.isfile(ROOT_DIR + 'LICENSE.md')
        RESULT = README and CHANGELOG and CODE_OF_CONDUCT and CONTRIBUTING and LICENSE
        return {
            "result": RESULT,
            "files": [
                {
                    "README.md": README
                },
                {
                    "CHANGELOG.md": CHANGELOG
                },
                {
                    "CODE_OF_CONDUCT.md": CODE_OF_CONDUCT
                },
                {
                    "CONTRIBUTING.md": CONTRIBUTING
                },
                {
                    "LICENSE.md": LICENSE
                }
            ]
        }

    @staticmethod
    def has_sonar():
        return os.path.isfile(ROOT_DIR + 'sonar-project.properties')

    @staticmethod
    def has_sonar_action():
        github_sonar_action = os.path.isfile(ROOT_DIR + '.github/workflows/sonar.yml') \
                              or os.path.isfile(ROOT_DIR + '.github/workflows/build.yml')
        return github_sonar_action

    @staticmethod
    def has_openapi():
        public_swagger = os.path.isfile(ROOT_DIR + 'public/swagger/index.html')
        # TODO incluir outros tipos de cenários aqui
        return public_swagger

    @staticmethod
    def has_restful():
        # TODO implementar futuramente
        return False

    @staticmethod
    def has_docker():
        return os.path.isdir(ROOT_DIR + 'docker')

    @staticmethod
    def has_editorconfig():
        return os.path.isfile(ROOT_DIR + '.editorconfig')

    @staticmethod
    def has_serverless():
        return os.path.isfile(ROOT_DIR + 'serverless.yaml')

    @staticmethod
    def has_log_monitoring():
        # TODO implementar futuramente
        return False

    @staticmethod
    def has_scripts():
        return os.path.isdir(ROOT_DIR + 'scripts')

    @staticmethod
    def has_environment_scripts():
        return os.path.isdir(ROOT_DIR + 'env')

    @staticmethod
    def has_terraform():
        return os.path.isdir(ROOT_DIR + 'infrastructure')

    @staticmethod
    def has_tests():
        if os.path.isdir(ROOT_DIR + 'test'):
            tests_folder = ROOT_DIR + 'test/'
        else:
            tests_folder = ROOT_DIR + 'tests/'

        ignored_dirs = ['.github', '.vscode', '.git', '.idea', 'docker', 'scripts', 'bin']
        directories = [f for f in os.listdir(ROOT_DIR) if
                       os.path.isdir(f) and f not in ignored_dirs]
        go_unit_pattern = r'(.*)_test\.go'
        node_unit_pattern = r'(.*)_spec\.(js|ts)'
        go_unit_test_files = True if GuidelinesChecker.search_by_pattern(directories, go_unit_pattern) else False
        node_unit_test_files = True if GuidelinesChecker.search_by_pattern(directories, node_unit_pattern) else False

        unit_tests = os.path.isdir(tests_folder + 'unit') \
                     or os.path.isdir(tests_folder + 'Unit') \
                     or os.path.isfile(tests_folder + 'jest-unit.json') \
                     or go_unit_test_files or node_unit_test_files
        integration_tests = os.path.isdir(tests_folder + 'integration') \
                            or os.path.isdir(tests_folder + 'Integration') \
                            or os.path.isfile(tests_folder + 'jest-integration.json')
        component_tests = os.path.isdir(tests_folder + 'component') \
                          or os.path.isdir(tests_folder + 'Component') \
                          or os.path.isfile(tests_folder + 'jest-component.json')
        result = unit_tests and integration_tests and component_tests

        return {
            "result": result,
            "contexts": [
                {
                    "unit": unit_tests
                },
                {
                    "integration": integration_tests
                },
                {
                    "component": component_tests
                },
            ]
        }

    @staticmethod
    def search_by_pattern(directories, go_pattern):
        test_files = []
        for d in directories:
            dir_files = []
            for f in os.listdir(os.path.join(ROOT_DIR, d)):
                if os.path.isdir(os.path.join(ROOT_DIR, d, f)):
                    for sf in os.listdir(os.path.join(ROOT_DIR, d, f)):
                        if re.match(go_pattern, str(sf)):
                            dir_files.append(sf)
                elif re.match(go_pattern, str(f)):
                    dir_files.append(f)
            #             print("dir_files", dir_files)
            test_files = test_files + dir_files
        return test_files

    @staticmethod
    def has_aws_ci_cd():
        buildspec = os.path.isfile(ROOT_DIR + 'buildspec.yaml')
        appspec = os.path.isfile(ROOT_DIR + 'appspec.yaml')

        return buildspec or appspec

    @staticmethod
    def has_meta_files():
        projectrc = os.path.isfile(ROOT_DIR + '.projectrc')
        dockerignore = os.path.isfile(ROOT_DIR + '.dockerignore')
        gitignore = os.path.isfile(ROOT_DIR + '.gitignore')
        docker_compose = os.path.isfile(ROOT_DIR + 'docker-compose.yml')

        return [
            {
                ".projectrc": projectrc
            },
            {
                ".dockerignore": dockerignore
            },
            {
                ".gitignore": gitignore
            },
            {
                "docker-compose.yml": docker_compose
            }
        ]


checker = GuidelinesChecker()
checker.check()

_LINE_SIZE = 80
print(repeat('-', 1, _LINE_SIZE))
print('{} - {}'.format('Guideline Checker', _VERSION_STR))
print(repeat('-', 1, _LINE_SIZE))
print('{}{}{}'.format('Guidelines'.ljust(40), 'Achieved/Found'.ljust(20), 'Required'))
print(repeat('-', 1, _LINE_SIZE))
for k in _GD_KEYS:
    value = _GUIDELINES[k]
    if isinstance(value, bool):
        print('{}{}{}'.format(k.ljust(40), str(value).ljust(20),
            GuidelinesChecker.GUIDES_CONFIG[k]['required']))
    elif isinstance(value, dict) or isinstance(value, list):
        print('{}{}{}'.format(k.ljust(40), ''.ljust(20),
            GuidelinesChecker.GUIDES_CONFIG[k]['required']))
        if isinstance(value, dict):
            sub_keys = value.keys()
        else:
            sub_keys = [i for i in range(0, len(value))]
        for sk in sub_keys:
            sv = value[sk]
            if isinstance(sv, list):
                print('{}{}{}'.format(''.ljust(5), str(sk).ljust(35), ''))
                for ssv in sv:
                    if isinstance(ssv, dict):
                        ssv_keys = list(ssv.keys())
                        key = ssv_keys[0]
                        val = ssv[key]
                        print('{}{}{}'.format(''.ljust(10), str(key).ljust(30), val))
                    else:
                        print('{}{}{}'.format(''.ljust(10), str(ssv).ljust(30), ssv))
            elif isinstance(sv, dict):
                sv_keys = list(sv.keys())
                key = sv_keys[0]
                val = sv[key]
                print('{}{}{}'.format(''.ljust(10), str(key).ljust(30), val))
            else:
                print('{}{}{}'.format(''.ljust(5), str(sk).ljust(35), sv))
print(repeat('-', 1, _LINE_SIZE))
print('{}{}'.format('Final Result'.ljust(40), _FINAL_RESULT))
print(repeat('-', 1, _LINE_SIZE))
# print('Generating the guidelines file: guidelines.yaml')

if _FINAL_RESULT:
    sys.exit(0)
else:
    sys.exit(1)
