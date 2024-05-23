namespace System.Collections.Generic
{
    using System.Linq;

    /// <summary>
    /// 
    /// </summary>
    public static class IEnumerableExtensions
    {
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="reference"></param>
        /// <returns></returns>
        public static bool IsNullOrEmpty<T>(this IEnumerable<T> reference)
        {
            return reference.IsNullOrEmpty() || !reference.Any();
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="reference"></param>
        /// <param name="rightLength"></param>
        /// <returns></returns>
        public static bool IsNullOrWrongLength<T>(this IEnumerable<T> reference, int rightLength = 0)
        {
            return reference == null || rightLength != reference.Count();
        }


        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, Action<int> body)
        {
            for (; fromInclusive < collection.Count();) body(fromInclusive++);
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, Action<T, int> body)
        {
            for (; fromInclusive < collection.Count();) body(collection.ElementAt(fromInclusive), fromInclusive++);
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, Func<T, bool> body)
        {
            for (; fromInclusive < collection.Count();) if (!body(collection.ElementAt(fromInclusive++))) break;
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, Func<T, int, bool> body)
        {
            for (; fromInclusive < collection.Count();) if (!body(collection.ElementAt(fromInclusive), fromInclusive++)) break;
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end with thread-local data in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <exception cref="AggregateException">The exception that contains all the individual exceptions thrown on all threads.</exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void For<T0, T1>(this IEnumerable<T0> collection, int fromInclusive, Func<T1> initial, Func<T0, T1, T1> body, Action<T1> final)
        {
            for (; fromInclusive < collection.Count();) final(body(collection.ElementAt(fromInclusive++), initial()));
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop to the special collection end with thread-local data in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <exception cref="AggregateException">The exception that contains all the individual exceptions thrown on all threads.</exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void For<T0, T1>(this IEnumerable<T0> collection, int fromInclusive, Func<T1> initial, Func<T0, int, T1, T1> body, Action<T1> final)
        {
            for (; fromInclusive < collection.Count();) final(body(collection.ElementAt(fromInclusive), fromInclusive++, initial()));
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, int toExclusive, Action<int> body)
        {
            for (; fromInclusive < toExclusive;) body(fromInclusive++);
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop with 32-bit indexes in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, int toExclusive, Action<T, int> body)
        {
            for (; fromInclusive < toExclusive;) body(collection.ElementAt(fromInclusive), fromInclusive++);
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, int toExclusive, Func<T, bool> body)
        {
            for (; fromInclusive < toExclusive;) if (!body(collection.ElementAt(fromInclusive++))) break;
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop with 32-bit indexes in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void For<T>(this IEnumerable<T> collection, int fromInclusive, int toExclusive, Func<T, int, bool> body)
        {
            for (; fromInclusive < toExclusive;) if (!body(collection.ElementAt(fromInclusive), fromInclusive++)) break;
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop with thread-local data in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <exception cref="AggregateException">The exception that contains all the individual exceptions thrown on all threads.</exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void For<T0, T1>(this IEnumerable<T0> collection, int fromInclusive, int toExclusive, Func<T1> initial, Func<T0, T1, T1> body, Action<T1> final)
        {
            for (; fromInclusive < toExclusive;) final(body(collection.ElementAt(fromInclusive++), initial()));
        }
        /// <summary>
        /// Executes a for (For in Visual Basic) loop with thread-local data in which iterations may run in sequential and the state of the loop can be manipulated.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <exception cref="AggregateException">The exception that contains all the individual exceptions thrown on all threads.</exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="fromInclusive">The start index, inclusive.</param>
        /// <param name="toExclusive">The end index, exclusive.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void For<T0, T1>(this IEnumerable<T0> collection, int fromInclusive, int toExclusive, Func<T1> initial, Func<T0, int, T1, T1> body, Action<T1> final)
        {
            for (; fromInclusive < toExclusive;) final(body(collection.ElementAt(fromInclusive), fromInclusive++, initial()));
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="TInput"></typeparam>
        /// <typeparam name="TOutput"></typeparam>
        /// <param name="collection"></param>
        /// <param name="fromInclusive"></param>
        /// <param name="convertor"></param>
        /// <returns></returns>
        public static IEnumerable<TOutput> For<TInput, TOutput>(this IEnumerable<TInput> collection, int fromInclusive, Func<TInput, TOutput> convertor)
        {
            var toExclusive = collection.Count();
            return For(collection, fromInclusive, toExclusive, convertor);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="TInput"></typeparam>
        /// <typeparam name="TOutput"></typeparam>
        /// <param name="collection"></param>
        /// <param name="fromInclusive"></param>
        /// <param name="toExclusive"></param>
        /// <param name="convertor"></param>
        /// <returns></returns>
        public static IEnumerable<TOutput> For<TInput, TOutput>(this IEnumerable<TInput> collection, int fromInclusive, int toExclusive, Func<TInput, TOutput> convertor)
        {
            for (; fromInclusive < toExclusive;)
                yield return convertor(collection.ElementAt(fromInclusive++));
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="TInput"></typeparam>
        /// <typeparam name="TOutput"></typeparam>
        /// <param name="collection"></param>
        /// <param name="fromInclusive"></param>
        /// <param name="convertor"></param>
        /// <returns></returns>
        public static IEnumerable<TOutput> For<TInput, TOutput>(this IEnumerable<TInput> collection, int fromInclusive, Func<TInput, int, TOutput> convertor)
        {
            var toExclusive = collection.Count();
            return For(collection, fromInclusive, toExclusive, convertor);
        }
        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="TInput"></typeparam>
        /// <typeparam name="TOutput"></typeparam>
        /// <param name="collection"></param>
        /// <param name="fromInclusive"></param>
        /// <param name="toExclusive"></param>
        /// <param name="convertor"></param>
        /// <returns></returns>
        public static IEnumerable<TOutput> For<TInput, TOutput>(this IEnumerable<TInput> collection, int fromInclusive, int toExclusive, Func<TInput, int, TOutput> convertor)
        {
            for (; fromInclusive < toExclusive;)
                yield return convertor(collection.ElementAt(fromInclusive), fromInclusive++);
        }

        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void ForEach<T>(this IEnumerable<T> collection, Action<T> body)
        {
            foreach (T instance in collection) body(instance);
        }
        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void ForEach<T>(this IEnumerable<T> collection, Action<T, int> body)
        {
            int index = 0;
            foreach (T instance in collection) body(instance, index++);
        }
        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void ForEach<T>(this IEnumerable<T> collection, Func<T, bool> body)
        {
            foreach (T instance in collection) if (!body(instance)) break;
        }
        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException">The collection argument is null.</exception>
        /// <typeparam name="T">The type of the data in the collection.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        public static void ForEach<T>(this IEnumerable<T> collection, Func<T, int, bool> body)
        {
            int index = 0;
            foreach (T instance in collection) if (!body(instance, index++)) break;
        }
        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation with thread-local data in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void ForEach<T0, T1>(this IEnumerable<T0> collection, Func<T1> initial, Func<T0, T1, T1> body, Action<T1> final)
        {
            foreach (T0 instance in collection) final(body(instance, initial()));
        }
        /// <summary>
        /// Executes a foreach (For Each in Visual Basic) operation with thread-local data in which iterations may run in sequential.
        /// </summary>
        /// <exception cref="ArgumentNullException"><para>The collection argument is null.</para><para>-or-</para><para>The body argument is null.</para><para>-or-</para><para>The initial argument is null.</para><para>-or-</para><para>The final argument is null.</para></exception>
        /// <typeparam name="T0">The type of the data in the collection.</typeparam>
        /// <typeparam name="T1">The type of the thread-local data.</typeparam>
        /// <param name="collection">An enumerable data collection.</param>
        /// <param name="initial">The function delegate that returns the initial state of the local data for each task.</param>
        /// <param name="body">The delegate that is invoked once per iteration.</param>
        /// <param name="final">The delegate that performs a final action on the local state of each task.</param>
        public static void ForEach<T0, T1>(this IEnumerable<T0> collection, Func<T1> initial, Func<T0, int, T1, T1> body, Action<T1> final)
        {
            int index = 0;
            foreach (T0 instance in collection) final(body(instance, index++, initial()));
        }
    }
}
