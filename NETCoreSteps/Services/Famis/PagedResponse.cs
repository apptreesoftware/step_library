using System;
using System.Collections;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace Famis
{
    public class PagedResponse<T> : IEnumerator<T>, IEnumerable<T>
    {
        public Uri NextLink { get; }

        public List<T> PageResults { get; }

        private readonly IService _client;

        public PagedResponse(List<T> pageResults, Uri nextLink, IService client) {
            PageResults = pageResults;
            _client = client;
            NextLink = nextLink;
        }

        public Task<PagedResponse<T>> NextPage() {
            if (!HasNextPage) {
                throw new Exception("NextPage not available. Check HasNextPage before calling this method");
            }
            return _client.GetNext(this);
        }

        public async Task<List<T>> GetAll() {
            var values = new List<T>(PageResults);
            var page = this;
            while(page.HasNextPage) {
                page = await page.NextPage();
                values.AddRange(page.PageResults);
            }
            return values;
        }
        
        public bool HasNextPage => NextLink != null;

        public int ResultCount => PageResults.Count;

        public bool MoveNext() {
            return GetEnumerator().MoveNext();
        }

        public void Reset() {
            GetEnumerator().Reset();
        }

        public T Current => GetEnumerator().Current;

        object IEnumerator.Current => GetEnumerator().Current;

        public void Dispose() {
            GetEnumerator().Dispose();
        }

        public IEnumerator<T> GetEnumerator() {
            return PageResults.GetEnumerator();
        }

        IEnumerator IEnumerable.GetEnumerator() {
            return GetEnumerator();
        }

        public void ForEach(Action<T> action) {
            PageResults.ForEach(action);
        }
        
        
    }
}