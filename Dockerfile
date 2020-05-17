ARG BASE=37-light
FROM elek/base:${BASE}
ARG ARTIFACTDIR
ADD ${ARTIFACTDIR}/weed /usr/bin/weed
RUN chmod +x /usr/bin/weed
