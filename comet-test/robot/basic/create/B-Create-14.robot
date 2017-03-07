*** Settings ***
Documentation          Basic test case (Create 30 volumes with API and create 6 containers with these volumes).
...                    

Library                OperatingSystem
Library                SSHLibrary
Suite Setup            Open Connection And Log In
Suite Teardown         Close All Connections

*** Variables ***
${HOST}                192.168.14.41
${USERNAME}            root
${PASSWORD}            123456
${APIHOST}             192.168.14.200
${PORT}                9876
${TESTCASE}            B-Create-14
${WORKDIR}             /root/dev
${LOOP}                6
${ID}                  0
${VOL_ID}              ${TESTCASE}-${ID}
${CAP}                 100

*** Test Cases ***
Create 30 volumes via API
    : FOR    ${INDEX}    IN RANGE    0     ${LOOP}
    \      Inner Loop1      ${INDEX}

Inspect 30 volumes via API
    : FOR    ${INDEX}    IN RANGE    0     ${LOOP}
    \      Inner Loop1-1      ${INDEX}

Start deploy onto host
    Log to console    ${HOST}
    ${output}=    Execute Command    echo Hello SSHLibrary!
    Should Be Equal    ${output}    Hello SSHLibrary!

Create 6 containers with 30 volumes
    : FOR    ${INDEX}    IN RANGE    0     ${LOOP}
    \     ${output}     ${stderr}=    Execute Command      bash ${WORKDIR}/comet-test/robot/basic/scripts/CreateFive.sh ${TESTCASE} ${INDEX} ${APIHOST}:${PORT}      return_stderr=True
    \     Log to console	\nContainer created with ID: ${output} ${stderr}
    \     Should Not Contain  ${output}    Robot-Fail

Delete container with docker command
    : FOR    ${INDEX}    IN RANGE    0      ${LOOP}
    \       ${output}    ${stderr}=    Execute Command      bash ${WORKDIR}/comet-test/robot/basic/scripts/DeleteFive.sh ${TESTCASE} ${INDEX} ${APIHOST}:${PORT}     return_stderr=True
    \       Log to console      \nContainer created with ID: ${output}    ${stderr}
    \       Should Not Contain  ${output}    Robot-Fail

Delete container with API again
    : FOR    ${INDEX}    IN RANGE    0      ${LOOP}
    \      Inner Loop2      ${INDEX}

Check again with API
    : FOR    ${INDEX}    IN RANGE    0      ${LOOP}
    \      Inner Loop3      ${INDEX}

*** Keywords ***
Open Connection And Log In
   Open Connection    ${HOST}
   Login    ${USERNAME}    ${PASSWORD}

Inner Loop1
    [Arguments]        ${Arg1}
    : FOR    ${I}    IN RANGE    0      5
    \      ${VOL_ID}=      Catenate     SEPARATOR=-    ${TESTCASE}-${ARG1}     ${I}
    \      ${CAP} =      Evaluate      	${I}*10 + 100
    \      ${output}=    Run      bash ${WORKDIR}/comet-test/robot/basic/scripts/api.sh ${APIHOST}:${PORT} create ${VOL_ID} ${CAP}
    \      Log to console	\n${output}
    \      Should Contain      ${output}      "Result": "0"

Inner Loop1-1
    [Arguments]        ${Arg1}
    : FOR    ${I}    IN RANGE    0      5
    \      ${VOL_ID}=      Catenate     SEPARATOR=-    ${TESTCASE}-${ARG1}     ${I}
    \      ${CAP} =      Evaluate      	${I}*10 + 100
    \      ${output}=    Run      bash ${WORKDIR}/comet-test/robot/basic/scripts/api.sh ${APIHOST}:${PORT} inspect ${VOL_ID}
    \      Log to console	\n${output}
    \      Should Contain      ${output}      "Result": "0"
    \      Should Contain      ${output}      "ID": "${VOL_ID}"
    \      Should Contain      ${output}      "Capacity": "${CAP}"

Inner Loop2
    [Arguments]        ${Arg1}
    : FOR    ${I}    IN RANGE    0      5
    \      ${VOL_ID}=      Catenate     SEPARATOR=-    ${TESTCASE}-${ARG1}     ${I}
    \      ${output}=    Run      bash ${WORKDIR}/comet-test/robot/basic/scripts/api.sh ${APIHOST}:${PORT} delete ${VOL_ID}
    \      Log to console       \n${output}
    \      Should Contain      ${output}      "Result": "0"

Inner Loop3
    [Arguments]        ${Arg1}
    : FOR    ${I}    IN RANGE    0      5
    \      ${VOL_ID}=      Catenate     SEPARATOR=-    ${TESTCASE}-${ARG1}     ${I}
    \      ${output}=    Run      bash ${WORKDIR}/comet-test/robot/basic/scripts/api.sh ${APIHOST}:${PORT} inspect ${VOL_ID}
    \      Log to console       \n${output}
    \      Should Contain      ${output}      "Result": "4000"
